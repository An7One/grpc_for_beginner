using System;
using System.Linq;
using System.Threading.Tasks;
using Grpc.Core;
using GrpcServer.Web.Data;
using GrpcServer.Web.Protos;
using Microsoft.Extensions.Logging;

namespace GrpcServer.Web
{
    public class MyEmployeeService : EmployeeService.EmployeeServiceBase
    {
        private readonly ILogger<MyEmployeeService> _logger;
        public MyEmployeeService(ILogger<MyEmployeeService> logger)
        {
            _logger = logger;
        }

        public override Task<EmployeeResponse> GetByNo(GetByNoRequest request, ServerCallContext context){
            var md = context.RequestHeaders;
            foreach(var pair in md){
                _logger.LogInformation($"{pair.Key}: {pair.Value}");
            }

            var employee = InMemoryData.Employees
                .SingleOrDefault(x => x.No == request.No);

            if(employee != null){
                var response = new EmployeeResponse{
                    Employee = employee
                };

                return Task.FromResult(response);
            }

            throw new Exception($"Employee not found with No: {request.No}");
        }
    }
}
