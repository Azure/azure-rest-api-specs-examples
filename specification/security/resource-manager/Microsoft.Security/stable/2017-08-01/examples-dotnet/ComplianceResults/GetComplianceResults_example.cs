using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2017-08-01/examples/ComplianceResults/GetComplianceResults_example.json
// this example is just showing the usage of "ComplianceResults_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ComplianceResultResource created on azure
// for more information of creating ComplianceResultResource, please refer to the document of ComplianceResultResource
string resourceId = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string complianceResultName = "DesignateMoreThanOneOwner";
ResourceIdentifier complianceResultResourceId = ComplianceResultResource.CreateResourceIdentifier(resourceId, complianceResultName);
ComplianceResultResource complianceResult = client.GetComplianceResultResource(complianceResultResourceId);

// invoke the operation
ComplianceResultResource result = await complianceResult.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ComplianceResultData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
