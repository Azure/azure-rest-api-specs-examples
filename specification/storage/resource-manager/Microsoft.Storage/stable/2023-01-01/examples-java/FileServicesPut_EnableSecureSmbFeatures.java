import com.azure.resourcemanager.storage.fluent.models.FileServicePropertiesInner;
import com.azure.resourcemanager.storage.models.ProtocolSettings;
import com.azure.resourcemanager.storage.models.SmbSetting;

/** Samples for FileServices SetServiceProperties. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/FileServicesPut_EnableSecureSmbFeatures.json
     */
    /**
     * Sample code: PutFileServices_EnableSecureSmbFeatures.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void putFileServicesEnableSecureSmbFeatures(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getFileServices()
            .setServicePropertiesWithResponse(
                "res4410",
                "sto8607",
                new FileServicePropertiesInner()
                    .withProtocolSettings(
                        new ProtocolSettings()
                            .withSmb(
                                new SmbSetting()
                                    .withVersions("SMB2.1;SMB3.0;SMB3.1.1")
                                    .withAuthenticationMethods("NTLMv2;Kerberos")
                                    .withKerberosTicketEncryption("RC4-HMAC;AES-256")
                                    .withChannelEncryption("AES-128-CCM;AES-128-GCM;AES-256-GCM"))),
                com.azure.core.util.Context.NONE);
    }
}
