using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;
using Azure.ResourceManager.DevCenter.Models;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/DevBoxDefinitions_Patch.json
// this example is just showing the usage of "DevBoxDefinitions_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevBoxDefinitionResource created on azure
// for more information of creating DevBoxDefinitionResource, please refer to the document of DevBoxDefinitionResource
string subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
string resourceGroupName = "rg1";
string devCenterName = "Contoso";
string devBoxDefinitionName = "WebDevBox";
ResourceIdentifier devBoxDefinitionResourceId = DevBoxDefinitionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, devCenterName, devBoxDefinitionName);
DevBoxDefinitionResource devBoxDefinition = client.GetDevBoxDefinitionResource(devBoxDefinitionResourceId);

// invoke the operation
DevBoxDefinitionPatch patch = new DevBoxDefinitionPatch()
{
    ImageReference = new DevCenterImageReference()
    {
        Id = new ResourceIdentifier("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/Example/providers/Microsoft.DevCenter/devcenters/Contoso/galleries/contosogallery/images/exampleImage/version/2.0.0"),
    },
};
ArmOperation<DevBoxDefinitionResource> lro = await devBoxDefinition.UpdateAsync(WaitUntil.Completed, patch);
DevBoxDefinitionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevBoxDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
