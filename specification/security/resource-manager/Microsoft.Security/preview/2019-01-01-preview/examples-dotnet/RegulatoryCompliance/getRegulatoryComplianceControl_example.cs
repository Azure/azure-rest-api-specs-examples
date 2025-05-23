using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/RegulatoryCompliance/getRegulatoryComplianceControl_example.json
// this example is just showing the usage of "RegulatoryComplianceControls_Get" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
RegulatoryComplianceControlResource result = await regulatoryComplianceControl.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RegulatoryComplianceControlData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
