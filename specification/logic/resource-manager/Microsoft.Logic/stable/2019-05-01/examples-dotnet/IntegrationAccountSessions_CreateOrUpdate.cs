using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Logic;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountSessions_CreateOrUpdate.json
// this example is just showing the usage of "IntegrationAccountSessions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IntegrationAccountResource created on azure
// for more information of creating IntegrationAccountResource, please refer to the document of IntegrationAccountResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string integrationAccountName = "testia123";
ResourceIdentifier integrationAccountResourceId = IntegrationAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, integrationAccountName);
IntegrationAccountResource integrationAccount = client.GetIntegrationAccountResource(integrationAccountResourceId);

// get the collection of this IntegrationAccountSessionResource
IntegrationAccountSessionCollection collection = integrationAccount.GetIntegrationAccountSessions();

// invoke the operation
string sessionName = "testsession123-ICN";
IntegrationAccountSessionData data = new IntegrationAccountSessionData(new AzureLocation("placeholder"))
{
    Content = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
        ["controlNumber"] = "1234",
        ["controlNumberChangedTime"] = "2017-02-21T22:30:11.9923759Z"
    }),
};
ArmOperation<IntegrationAccountSessionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, sessionName, data);
IntegrationAccountSessionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IntegrationAccountSessionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
