using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/RegulatoryCompliance/getRegulatoryComplianceControlList_example.json
// this example is just showing the usage of "RegulatoryComplianceControls_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RegulatoryComplianceStandardResource created on azure
// for more information of creating RegulatoryComplianceStandardResource, please refer to the document of RegulatoryComplianceStandardResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string regulatoryComplianceStandardName = "PCI-DSS-3.2";
ResourceIdentifier regulatoryComplianceStandardResourceId = RegulatoryComplianceStandardResource.CreateResourceIdentifier(subscriptionId, regulatoryComplianceStandardName);
RegulatoryComplianceStandardResource regulatoryComplianceStandard = client.GetRegulatoryComplianceStandardResource(regulatoryComplianceStandardResourceId);

// get the collection of this RegulatoryComplianceControlResource
RegulatoryComplianceControlCollection collection = regulatoryComplianceStandard.GetRegulatoryComplianceControls();

// invoke the operation and iterate over the result
await foreach (RegulatoryComplianceControlResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    RegulatoryComplianceControlData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
