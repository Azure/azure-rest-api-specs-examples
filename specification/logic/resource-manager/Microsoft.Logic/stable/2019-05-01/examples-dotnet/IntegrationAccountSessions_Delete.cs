using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountSessions_Delete.json
// this example is just showing the usage of "IntegrationAccountSessions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IntegrationAccountSessionResource created on azure
// for more information of creating IntegrationAccountSessionResource, please refer to the document of IntegrationAccountSessionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string integrationAccountName = "testia123";
string sessionName = "testsession123-ICN";
ResourceIdentifier integrationAccountSessionResourceId = IntegrationAccountSessionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, integrationAccountName, sessionName);
IntegrationAccountSessionResource integrationAccountSession = client.GetIntegrationAccountSessionResource(integrationAccountSessionResourceId);

// invoke the operation
await integrationAccountSession.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
