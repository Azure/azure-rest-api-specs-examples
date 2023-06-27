using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PartnerNamespaces_Update.json
// this example is just showing the usage of "PartnerNamespaces_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PartnerNamespaceResource created on azure
// for more information of creating PartnerNamespaceResource, please refer to the document of PartnerNamespaceResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string partnerNamespaceName = "examplePartnerNamespaceName1";
ResourceIdentifier partnerNamespaceResourceId = PartnerNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, partnerNamespaceName);
PartnerNamespaceResource partnerNamespace = client.GetPartnerNamespaceResource(partnerNamespaceResourceId);

// invoke the operation
PartnerNamespacePatch patch = new PartnerNamespacePatch()
{
    Tags =
    {
    ["tag1"] = "value1",
    },
};
await partnerNamespace.UpdateAsync(WaitUntil.Completed, patch);

Console.WriteLine($"Succeeded");
