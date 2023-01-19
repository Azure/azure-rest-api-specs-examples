using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;
using Azure.ResourceManager.Logic.Models;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountPartners_Delete.json
// this example is just showing the usage of "IntegrationAccountPartners_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IntegrationAccountPartnerResource created on azure
// for more information of creating IntegrationAccountPartnerResource, please refer to the document of IntegrationAccountPartnerResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string integrationAccountName = "testIntegrationAccount";
string partnerName = "testPartner";
ResourceIdentifier integrationAccountPartnerResourceId = IntegrationAccountPartnerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, integrationAccountName, partnerName);
IntegrationAccountPartnerResource integrationAccountPartner = client.GetIntegrationAccountPartnerResource(integrationAccountPartnerResourceId);

// invoke the operation
await integrationAccountPartner.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
