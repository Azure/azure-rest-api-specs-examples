using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedServices.Models;
using Azure.ResourceManager.ManagedServices;

// Generated from example definition: specification/managedservices/resource-manager/Microsoft.ManagedServices/stable/2022-10-01/examples/GetRegistrationAssignment.json
// this example is just showing the usage of "RegistrationAssignments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this ManagedServicesRegistrationAssignmentResource
string scope = "subscription/0afefe50-734e-4610-8a82-a144ahf49dea";
ManagedServicesRegistrationAssignmentCollection collection = client.GetManagedServicesRegistrationAssignments(new ResourceIdentifier(scope));

// invoke the operation
string registrationAssignmentId = "26c128c2-fefa-4340-9bb1-6e081c90ada2";
NullableResponse<ManagedServicesRegistrationAssignmentResource> response = await collection.GetIfExistsAsync(registrationAssignmentId);
ManagedServicesRegistrationAssignmentResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ManagedServicesRegistrationAssignmentData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
