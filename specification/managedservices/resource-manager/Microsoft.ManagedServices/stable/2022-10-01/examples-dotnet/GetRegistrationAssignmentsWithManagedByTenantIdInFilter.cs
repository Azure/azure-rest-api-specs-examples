using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedServices.Models;
using Azure.ResourceManager.ManagedServices;

// Generated from example definition: specification/managedservices/resource-manager/Microsoft.ManagedServices/stable/2022-10-01/examples/GetRegistrationAssignmentsWithManagedByTenantIdInFilter.json
// this example is just showing the usage of "RegistrationAssignments_List" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this ManagedServicesRegistrationAssignmentResource
string scope = "subscription/0afefe50-734e-4610-8a82-a144ahf49dea";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
ManagedServicesRegistrationAssignmentCollection collection = client.GetManagedServicesRegistrationAssignments(scopeId);

// invoke the operation and iterate over the result
string filter = "$filter=managedByTenantId in (83abe5cd-bcc3-441a-bd86-e6a75360cec, de83f4a9-a76a-4025-a91a-91171923eac7)";
await foreach (ManagedServicesRegistrationAssignmentResource item in collection.GetAllAsync(filter: filter))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ManagedServicesRegistrationAssignmentData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
