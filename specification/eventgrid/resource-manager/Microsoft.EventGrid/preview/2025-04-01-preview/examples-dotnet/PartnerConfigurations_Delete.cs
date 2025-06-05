using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/PartnerConfigurations_Delete.json
// this example is just showing the usage of "PartnerConfigurations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PartnerConfigurationResource created on azure
// for more information of creating PartnerConfigurationResource, please refer to the document of PartnerConfigurationResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
ResourceIdentifier partnerConfigurationResourceId = PartnerConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
PartnerConfigurationResource partnerConfiguration = client.GetPartnerConfigurationResource(partnerConfigurationResourceId);

// invoke the operation
await partnerConfiguration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
