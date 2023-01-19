using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;
using Azure.ResourceManager.Logic.Models;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAgreements_List.json
// this example is just showing the usage of "IntegrationAccountAgreements_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IntegrationAccountResource created on azure
// for more information of creating IntegrationAccountResource, please refer to the document of IntegrationAccountResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string integrationAccountName = "testIntegrationAccount";
ResourceIdentifier integrationAccountResourceId = IntegrationAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, integrationAccountName);
IntegrationAccountResource integrationAccount = client.GetIntegrationAccountResource(integrationAccountResourceId);

// get the collection of this IntegrationAccountAgreementResource
IntegrationAccountAgreementCollection collection = integrationAccount.GetIntegrationAccountAgreements();

// invoke the operation and iterate over the result
await foreach (IntegrationAccountAgreementResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    IntegrationAccountAgreementData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
