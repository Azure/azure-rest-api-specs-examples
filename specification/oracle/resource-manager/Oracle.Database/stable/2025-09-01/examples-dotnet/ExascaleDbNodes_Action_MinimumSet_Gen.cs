using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OracleDatabase.Models;
using Azure.ResourceManager.OracleDatabase;

// Generated from example definition: 2025-09-01/ExascaleDbNodes_Action_MinimumSet_Gen.json
// this example is just showing the usage of "ExascaleDbNodes_Action" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExascaleDBNodeResource created on azure
// for more information of creating ExascaleDBNodeResource, please refer to the document of ExascaleDBNodeResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rgopenapi";
string exadbVmClusterName = "exadbvmcluster1";
string exascaleDbNodeName = "exascaledbnode1";
ResourceIdentifier exascaleDBNodeResourceId = ExascaleDBNodeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, exadbVmClusterName, exascaleDbNodeName);
ExascaleDBNodeResource exascaleDBNode = client.GetExascaleDBNodeResource(exascaleDBNodeResourceId);

// invoke the operation
DBNodeAction body = new DBNodeAction(DBNodeActionType.Start);
ArmOperation<ExascaleDBNodeActionResult> lro = await exascaleDBNode.ActionAsync(WaitUntil.Completed, body);
ExascaleDBNodeActionResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
