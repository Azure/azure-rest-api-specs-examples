using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OracleDatabase.Models;
using Azure.ResourceManager.OracleDatabase;

// Generated from example definition: 2025-03-01/oracleSubscriptions_listActivationLinks.json
// this example is just showing the usage of "OracleSubscriptions_ListActivationLinks" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OracleSubscriptionResource created on azure
// for more information of creating OracleSubscriptionResource, please refer to the document of OracleSubscriptionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
ResourceIdentifier oracleSubscriptionResourceId = OracleSubscriptionResource.CreateResourceIdentifier(subscriptionId);
OracleSubscriptionResource oracleSubscription = client.GetOracleSubscriptionResource(oracleSubscriptionResourceId);

// invoke the operation
ArmOperation<CloudAccountActivationLinks> lro = await oracleSubscription.GetActivationLinksAsync(WaitUntil.Completed);
CloudAccountActivationLinks result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
