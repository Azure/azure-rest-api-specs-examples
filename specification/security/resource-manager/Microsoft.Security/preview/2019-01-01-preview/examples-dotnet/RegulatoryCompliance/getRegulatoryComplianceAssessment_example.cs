using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/RegulatoryCompliance/getRegulatoryComplianceAssessment_example.json
// this example is just showing the usage of "RegulatoryComplianceAssessments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RegulatoryComplianceControlResource created on azure
// for more information of creating RegulatoryComplianceControlResource, please refer to the document of RegulatoryComplianceControlResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string regulatoryComplianceStandardName = "PCI-DSS-3.2";
string regulatoryComplianceControlName = "1.1";
ResourceIdentifier regulatoryComplianceControlResourceId = RegulatoryComplianceControlResource.CreateResourceIdentifier(subscriptionId, regulatoryComplianceStandardName, regulatoryComplianceControlName);
RegulatoryComplianceControlResource regulatoryComplianceControl = client.GetRegulatoryComplianceControlResource(regulatoryComplianceControlResourceId);

// get the collection of this RegulatoryComplianceAssessmentResource
RegulatoryComplianceAssessmentCollection collection = regulatoryComplianceControl.GetRegulatoryComplianceAssessments();

// invoke the operation
string regulatoryComplianceAssessmentName = "968548cb-02b3-8cd2-11f8-0cf64ab1a347";
bool result = await collection.ExistsAsync(regulatoryComplianceAssessmentName);

Console.WriteLine($"Succeeded: {result}");
