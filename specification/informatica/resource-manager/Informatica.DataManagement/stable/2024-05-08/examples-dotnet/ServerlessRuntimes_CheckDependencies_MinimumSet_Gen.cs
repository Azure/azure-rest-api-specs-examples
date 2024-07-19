using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.InformaticaDataManagement.Models;
using Azure.ResourceManager.InformaticaDataManagement;

// Generated from example definition: specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_CheckDependencies_MinimumSet_Gen.json
// this example is just showing the usage of "ServerlessRuntimes_CheckDependencies" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InformaticaServerlessRuntimeResource created on azure
// for more information of creating InformaticaServerlessRuntimeResource, please refer to the document of InformaticaServerlessRuntimeResource
string subscriptionId = "3599DA28-E346-4D9F-811E-189C0445F0FE";
string resourceGroupName = "rgopenapi";
string organizationName = "_-";
string serverlessRuntimeName = "_2_";
ResourceIdentifier informaticaServerlessRuntimeResourceId = InformaticaServerlessRuntimeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, organizationName, serverlessRuntimeName);
InformaticaServerlessRuntimeResource informaticaServerlessRuntime = client.GetInformaticaServerlessRuntimeResource(informaticaServerlessRuntimeResourceId);

// invoke the operation
CheckDependenciesResult result = await informaticaServerlessRuntime.CheckDependenciesAsync();

Console.WriteLine($"Succeeded: {result}");
