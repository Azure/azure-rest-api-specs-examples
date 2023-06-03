using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2017-08-01/examples/ComplianceResults/GetComplianceResults_example.json
// this example is just showing the usage of "ComplianceResults_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this ComplianceResultResource
string resourceId = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", resourceId));
ComplianceResultCollection collection = client.GetComplianceResults(scopeId);

// invoke the operation
string complianceResultName = "DesignateMoreThanOneOwner";
bool result = await collection.ExistsAsync(complianceResultName);

Console.WriteLine($"Succeeded: {result}");
