using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataBoxEdge;
using Azure.ResourceManager.DataBoxEdge.Models;

// Generated from example definition: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/RoleDelete.json
// this example is just showing the usage of "Roles_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataBoxEdgeRoleResource created on azure
// for more information of creating DataBoxEdgeRoleResource, please refer to the document of DataBoxEdgeRoleResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "GroupForEdgeAutomation";
string deviceName = "testedgedevice";
string name = "IoTRole1";
ResourceIdentifier dataBoxEdgeRoleResourceId = DataBoxEdgeRoleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deviceName, name);
DataBoxEdgeRoleResource dataBoxEdgeRole = client.GetDataBoxEdgeRoleResource(dataBoxEdgeRoleResourceId);

// invoke the operation
await dataBoxEdgeRole.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
