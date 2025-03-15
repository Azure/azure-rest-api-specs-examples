using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataBoxEdge.Models;
using Azure.ResourceManager.DataBoxEdge;

// Generated from example definition: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/PutAddons.json
// this example is just showing the usage of "Addons_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataBoxEdgeRoleAddonResource created on azure
// for more information of creating DataBoxEdgeRoleAddonResource, please refer to the document of DataBoxEdgeRoleAddonResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "GroupForEdgeAutomation";
string deviceName = "testedgedevice";
string roleName = "KubernetesRole";
string addonName = "arcName";
ResourceIdentifier dataBoxEdgeRoleAddonResourceId = DataBoxEdgeRoleAddonResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deviceName, roleName, addonName);
DataBoxEdgeRoleAddonResource dataBoxEdgeRoleAddon = client.GetDataBoxEdgeRoleAddonResource(dataBoxEdgeRoleAddonResourceId);

// invoke the operation
DataBoxEdgeRoleAddonData data = new EdgeArcAddon("4385cf00-2d3a-425a-832f-f4285b1c9dce", "GroupForEdgeAutomation", "testedgedevice", new AzureLocation("EastUS"));
ArmOperation<DataBoxEdgeRoleAddonResource> lro = await dataBoxEdgeRoleAddon.UpdateAsync(WaitUntil.Completed, data);
DataBoxEdgeRoleAddonResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataBoxEdgeRoleAddonData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
