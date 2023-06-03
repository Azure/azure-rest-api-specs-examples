using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/BackupResourceEncryptionConfig_Get.json
// this example is just showing the usage of "BackupResourceEncryptionConfigs_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupResourceEncryptionConfigExtendedResource created on azure
// for more information of creating BackupResourceEncryptionConfigExtendedResource, please refer to the document of BackupResourceEncryptionConfigExtendedResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rishgrp";
string vaultName = "rishTestVault";
ResourceIdentifier backupResourceEncryptionConfigExtendedResourceId = BackupResourceEncryptionConfigExtendedResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName);
BackupResourceEncryptionConfigExtendedResource backupResourceEncryptionConfigExtended = client.GetBackupResourceEncryptionConfigExtendedResource(backupResourceEncryptionConfigExtendedResourceId);

// invoke the operation
BackupResourceEncryptionConfigExtendedResource result = await backupResourceEncryptionConfigExtended.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BackupResourceEncryptionConfigExtendedData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
