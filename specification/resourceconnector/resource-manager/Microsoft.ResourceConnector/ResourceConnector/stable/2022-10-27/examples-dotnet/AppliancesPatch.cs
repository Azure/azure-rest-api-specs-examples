using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ResourceConnector.Models;
using Azure.ResourceManager.ResourceConnector;

// Generated from example definition: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/ResourceConnector/stable/2022-10-27/examples/AppliancesPatch.json
// this example is just showing the usage of "Appliances_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceConnectorApplianceResource created on azure
// for more information of creating ResourceConnectorApplianceResource, please refer to the document of ResourceConnectorApplianceResource
string subscriptionId = "11111111-2222-3333-4444-555555555555";
string resourceGroupName = "testresourcegroup";
string resourceName = "appliance01";
ResourceIdentifier resourceConnectorApplianceResourceId = ResourceConnectorApplianceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
ResourceConnectorApplianceResource resourceConnectorAppliance = client.GetResourceConnectorApplianceResource(resourceConnectorApplianceResourceId);

// invoke the operation
ResourceConnectorAppliancePatch patch = new ResourceConnectorAppliancePatch
{
    Tags =
    {
    ["key"] = "value"
    },
};
ResourceConnectorApplianceResource result = await resourceConnectorAppliance.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ResourceConnectorApplianceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
