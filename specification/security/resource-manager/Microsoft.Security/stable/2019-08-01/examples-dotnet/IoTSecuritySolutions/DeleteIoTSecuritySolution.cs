using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/DeleteIoTSecuritySolution.json
// this example is just showing the usage of "IotSecuritySolution_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotSecuritySolutionResource created on azure
// for more information of creating IotSecuritySolutionResource, please refer to the document of IotSecuritySolutionResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "MyGroup";
string solutionName = "default";
ResourceIdentifier iotSecuritySolutionResourceId = IotSecuritySolutionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, solutionName);
IotSecuritySolutionResource iotSecuritySolution = client.GetIotSecuritySolutionResource(iotSecuritySolutionResourceId);

// invoke the operation
await iotSecuritySolution.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
