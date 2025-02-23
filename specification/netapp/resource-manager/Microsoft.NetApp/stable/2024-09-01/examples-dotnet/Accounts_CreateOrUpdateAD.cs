using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-09-01/examples/Accounts_CreateOrUpdateAD.json
// this example is just showing the usage of "Accounts_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
string resourceGroupName = "myRG";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this NetAppAccountResource
NetAppAccountCollection collection = resourceGroupResource.GetNetAppAccounts();

// invoke the operation
string accountName = "account1";
NetAppAccountData data = new NetAppAccountData(new AzureLocation("eastus"))
{
    ActiveDirectories = {new NetAppAccountActiveDirectory
    {
    Username = "ad_user_name",
    Password = "ad_password",
    Domain = "10.10.10.3",
    Dns = "10.10.10.3",
    SmbServerName = "SMBServer",
    OrganizationalUnit = "OU=Engineering",
    Site = "SiteName",
    IsAesEncryptionEnabled = true,
    IsLdapSigningEnabled = false,
    IsLdapOverTlsEnabled = false,
    }},
};
ArmOperation<NetAppAccountResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, accountName, data);
NetAppAccountResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetAppAccountData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
