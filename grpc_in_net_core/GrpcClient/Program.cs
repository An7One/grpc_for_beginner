using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Grpc.Core;
using Grpc.Net.Client;
using GrpcClient.Protos;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Hosting;

namespace GrpcClient
{
    public class Program
    {
        public static async Task Main(string[] args)
        {
            // for mac os
            AppContext.SetSwitch("System.Net.Http.SocketsHttpHandler.Http2UnencryptedSupport", true);

            var md = new Metadata{
                {"username", "dave"},
                {"role", "adminstrator"},
            };

            using var channel = GrpcChannel.ForAddress("http://localhost:5001");
            var client = new EmployeeService.EmployeeServiceClient(channel);

            var response = await client.GetByNoAsync(new GetByNoRequest{
                No = 1994
            }, md);

            Console.WriteLine($"Response messages: {response}");
            Console.WriteLine("Press any key to exit.");
            Console.ReadKey();
        }
    }
}
