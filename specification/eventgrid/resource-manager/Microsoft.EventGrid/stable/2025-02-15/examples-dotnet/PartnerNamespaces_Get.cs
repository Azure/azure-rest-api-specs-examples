using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/PartnerNamespaces_Get.json
// this example is just showing the usage of "PartnerNamespaces_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PartnerNamespaceResource created on azure
// for more information of creating PartnerNamespaceResource, please refer to the document of PartnerNamespaceResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
string partnerNamespaceName = "examplePartnerNamespaceName1";
ResourceIdentifier partnerNamespaceResourceId = PartnerNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, partnerNamespaceName);
PartnerNamespaceResource partnerNamespace = client.GetPartnerNamespaceResource(partnerNamespaceResourceId);

// invoke the operation
PartnerNamespaceResource result = await partnerNamespace.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PartnerNamespaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
