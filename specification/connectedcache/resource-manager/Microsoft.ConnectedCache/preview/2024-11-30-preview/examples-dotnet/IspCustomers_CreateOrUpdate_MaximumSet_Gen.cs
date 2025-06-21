using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ConnectedCache.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ConnectedCache;

// Generated from example definition: 2024-11-30-preview/IspCustomers_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "IspCustomerResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "rgConnectedCache";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this IspCustomerResource
IspCustomerCollection collection = resourceGroupResource.GetIspCustomers();

// invoke the operation
string customerResourceName = "MccRPTest2";
IspCustomerData data = new IspCustomerData(new AzureLocation("westus"))
{
    Properties = new MccCustomerProperty
    {
        Customer = new MccCustomerEntity
        {
            FullyQualifiedResourceId = new ResourceIdentifier("uqsbtgae"),
            CustomerName = "mkpzynfqihnjfdbaqbqwyhd",
            ContactEmail = "xquos",
            ContactPhone = "vue",
            ContactName = "wxyqjoyoscmvimgwhpitxky",
            IsEntitled = true,
            ReleaseVersion = 20,
            ClientTenantId = "fproidkpgvpdnac",
            IsEnterpriseManaged = true,
            ShouldMigrate = true,
            ResendSignupCode = true,
            VerifySignupCode = true,
            VerifySignupPhrase = "tprjvttkgmrqlsyicnidhm",
        },
        AdditionalCustomerProperties = new MccCustomerAdditionalProperties
        {
            CustomerEmail = "zdjgibsidydyzm",
            CustomerTransitAsn = "habgklnxqzmozqpazoyejwiphezpi",
            CustomerTransitState = new CustomerTransitState("voblixkxfejbmhxilb"),
            CustomerAsn = "hgrelgnrtdkleisnepfolu",
            CustomerEntitlementSkuId = "b",
            CustomerEntitlementSkuGuid = "rvzmdpxyflgqetvpwupnfaxsweiiz",
            CustomerEntitlementSkuName = "waaqfijr",
            CustomerEntitlementExpiryOn = DateTimeOffset.Parse("2024-01-30T00:54:04.773Z"),
            OptionalProperty1 = "qhmwxza",
            OptionalProperty2 = "l",
            OptionalProperty3 = "mblwwvbie",
            OptionalProperty4 = "vzuek",
            OptionalProperty5 = "fzjodscdfcdr",
        },
    },
    Tags =
    {
    ["key1878"] = "warz"
    },
};
ArmOperation<IspCustomerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, customerResourceName, data);
IspCustomerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IspCustomerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
