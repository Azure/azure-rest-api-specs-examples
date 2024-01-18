using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2023-01-01-preview/examples/SecurityOperators/PutSecurityOperatorByName_example.json
// this example is just showing the usage of "SecurityOperators_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityOperatorResource created on azure
// for more information of creating SecurityOperatorResource, please refer to the document of SecurityOperatorResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string pricingName = "CloudPosture";
string securityOperatorName = "DefenderCSPMSecurityOperator";
ResourceIdentifier securityOperatorResourceId = SecurityOperatorResource.CreateResourceIdentifier(subscriptionId, pricingName, securityOperatorName);
SecurityOperatorResource securityOperator = client.GetSecurityOperatorResource(securityOperatorResourceId);

// invoke the operation
ArmOperation<SecurityOperatorResource> lro = await securityOperator.UpdateAsync(WaitUntil.Completed);
SecurityOperatorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityOperatorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
