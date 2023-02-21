using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ResumeSqlPool.json
// this example is just showing the usage of "SqlPools_Resume" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseSqlPoolResource created on azure
// for more information of creating SynapseSqlPoolResource, please refer to the document of SynapseSqlPoolResource
string subscriptionId = "01234567-89ab-4def-0123-456789abcdef";
string resourceGroupName = "sqlcrudtest-6852";
string workspaceName = "sqlcrudtest-2080";
string sqlPoolName = "sqlcrudtest-9187";
ResourceIdentifier synapseSqlPoolResourceId = SynapseSqlPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, sqlPoolName);
SynapseSqlPoolResource synapseSqlPool = client.GetSynapseSqlPoolResource(synapseSqlPoolResourceId);

// invoke the operation
ArmOperation<BinaryData> lro = await synapseSqlPool.ResumeAsync(WaitUntil.Completed);
BinaryData result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
