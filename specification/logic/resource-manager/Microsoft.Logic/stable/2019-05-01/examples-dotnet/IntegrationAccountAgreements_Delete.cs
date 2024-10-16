using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Logic.Models;
using Azure.ResourceManager.Logic;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAgreements_Delete.json
// this example is just showing the usage of "IntegrationAccountAgreements_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IntegrationAccountAgreementResource created on azure
// for more information of creating IntegrationAccountAgreementResource, please refer to the document of IntegrationAccountAgreementResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string integrationAccountName = "testIntegrationAccount";
string agreementName = "testAgreement";
ResourceIdentifier integrationAccountAgreementResourceId = IntegrationAccountAgreementResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, integrationAccountName, agreementName);
IntegrationAccountAgreementResource integrationAccountAgreement = client.GetIntegrationAccountAgreementResource(integrationAccountAgreementResourceId);

// invoke the operation
await integrationAccountAgreement.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
