using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ProviderHub.Models;
using Azure.ResourceManager.ProviderHub;

// Generated from example definition: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/AuthorizedApplications_CreateOrUpdate.json
// this example is just showing the usage of "AuthorizedApplications_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProviderAuthorizedApplicationResource created on azure
// for more information of creating ProviderAuthorizedApplicationResource, please refer to the document of ProviderAuthorizedApplicationResource
string subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
string providerNamespace = "Microsoft.Contoso";
Guid applicationId = Guid.Parse("760505bf-dcfa-4311-b890-18da392a00b2");
ResourceIdentifier providerAuthorizedApplicationResourceId = ProviderAuthorizedApplicationResource.CreateResourceIdentifier(subscriptionId, providerNamespace, applicationId);
ProviderAuthorizedApplicationResource providerAuthorizedApplication = client.GetProviderAuthorizedApplicationResource(providerAuthorizedApplicationResourceId);

// invoke the operation
ProviderAuthorizedApplicationData data = new ProviderAuthorizedApplicationData
{
    Properties = new ProviderAuthorizedApplicationProperties
    {
        ProviderAuthorization = new ApplicationProviderAuthorization
        {
            RoleDefinitionId = "123456bf-gkur-2098-b890-98da392a00b2",
            ManagedByRoleDefinitionId = "1a3b5c7d-8e9f-10g1-1h12-i13j14k1",
        },
        DataAuthorizations = {new ApplicationDataAuthorization(ApplicationOwnershipRole.ServiceOwner)
        {
        ResourceTypes = {"*"},
        }},
    },
};
ArmOperation<ProviderAuthorizedApplicationResource> lro = await providerAuthorizedApplication.UpdateAsync(WaitUntil.Completed, data);
ProviderAuthorizedApplicationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ProviderAuthorizedApplicationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
