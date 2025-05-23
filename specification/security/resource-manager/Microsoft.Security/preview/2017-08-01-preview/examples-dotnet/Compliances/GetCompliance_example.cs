using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/Compliances/GetCompliance_example.json
// this example is just showing the usage of "Compliances_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityComplianceResource created on azure
// for more information of creating SecurityComplianceResource, please refer to the document of SecurityComplianceResource
string scope = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string complianceName = "2018-01-01Z";
ResourceIdentifier securityComplianceResourceId = SecurityComplianceResource.CreateResourceIdentifier(scope, complianceName);
SecurityComplianceResource securityCompliance = client.GetSecurityComplianceResource(securityComplianceResourceId);

// invoke the operation
SecurityComplianceResource result = await securityCompliance.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityComplianceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
