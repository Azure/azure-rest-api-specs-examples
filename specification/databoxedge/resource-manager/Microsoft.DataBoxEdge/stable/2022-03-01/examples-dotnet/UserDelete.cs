using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataBoxEdge;
using Azure.ResourceManager.DataBoxEdge.Models;

// Generated from example definition: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/UserDelete.json
// this example is just showing the usage of "Users_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataBoxEdgeUserResource created on azure
// for more information of creating DataBoxEdgeUserResource, please refer to the document of DataBoxEdgeUserResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "GroupForEdgeAutomation";
string deviceName = "testedgedevice";
string name = "user1";
ResourceIdentifier dataBoxEdgeUserResourceId = DataBoxEdgeUserResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deviceName, name);
DataBoxEdgeUserResource dataBoxEdgeUser = client.GetDataBoxEdgeUserResource(dataBoxEdgeUserResourceId);

// invoke the operation
await dataBoxEdgeUser.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
