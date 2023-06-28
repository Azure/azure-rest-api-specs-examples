using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/PartnerConfigurations_UnauthorizePartner.json
// this example is just showing the usage of "PartnerConfigurations_UnauthorizePartner" operation, for the dependent resources, they will have to be created separately.

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
EventGridPartnerContent content = new EventGridPartnerContent()
{
    PartnerRegistrationImmutableId = Guid.Parse("941892bc-f5d0-4d1c-8fb5-477570fc2b71"),
    PartnerName = "Contoso.Finance",
    AuthorizationExpireOn = DateTimeOffset.Parse("2022-01-28T01:20:55.142Z"),
};
PartnerConfigurationResource result = await partnerConfiguration.UnauthorizePartnerAsync(content);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PartnerConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
