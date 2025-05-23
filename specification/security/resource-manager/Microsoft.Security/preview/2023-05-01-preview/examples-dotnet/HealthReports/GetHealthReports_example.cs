using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2023-05-01-preview/examples/HealthReports/GetHealthReports_example.json
// this example is just showing the usage of "HealthReports_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityHealthReportResource created on azure
// for more information of creating SecurityHealthReportResource, please refer to the document of SecurityHealthReportResource
string resourceId = "subscriptions/a1efb6ca-fbc5-4782-9aaa-5c7daded1ce2/resourcegroups/E2E-IBB0WX/providers/Microsoft.Security/securityconnectors/AwsConnectorAllOfferings";
string healthReportName = "909c629a-bf39-4521-8e4f-10b443a0bc02";
ResourceIdentifier securityHealthReportResourceId = SecurityHealthReportResource.CreateResourceIdentifier(resourceId, healthReportName);
SecurityHealthReportResource securityHealthReport = client.GetSecurityHealthReportResource(securityHealthReportResourceId);

// invoke the operation
SecurityHealthReportResource result = await securityHealthReport.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityHealthReportData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
