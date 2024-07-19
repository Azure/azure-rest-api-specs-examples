using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.InformaticaDataManagement.Models;
using Azure.ResourceManager.InformaticaDataManagement;

// Generated from example definition: specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_CreateOrUpdate_MinimumSet_Gen.json
// this example is just showing the usage of "ServerlessRuntimes_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InformaticaOrganizationResource created on azure
// for more information of creating InformaticaOrganizationResource, please refer to the document of InformaticaOrganizationResource
string subscriptionId = "3599DA28-E346-4D9F-811E-189C0445F0FE";
string resourceGroupName = "rgopenapi";
string organizationName = "-4Z__7";
ResourceIdentifier informaticaOrganizationResourceId = InformaticaOrganizationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, organizationName);
InformaticaOrganizationResource informaticaOrganization = client.GetInformaticaOrganizationResource(informaticaOrganizationResourceId);

// get the collection of this InformaticaServerlessRuntimeResource
InformaticaServerlessRuntimeCollection collection = informaticaOrganization.GetInformaticaServerlessRuntimes();

// invoke the operation
string serverlessRuntimeName = "J";
InformaticaServerlessRuntimeData data = new InformaticaServerlessRuntimeData();
ArmOperation<InformaticaServerlessRuntimeResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serverlessRuntimeName, data);
InformaticaServerlessRuntimeResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
InformaticaServerlessRuntimeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
