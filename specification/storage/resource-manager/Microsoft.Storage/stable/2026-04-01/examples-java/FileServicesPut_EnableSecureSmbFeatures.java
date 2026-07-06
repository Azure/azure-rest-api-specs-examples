
import com.azure.resourcemanager.storage.fluent.models.FileServicePropertiesInner;
import com.azure.resourcemanager.storage.models.ProtocolSettings;
import com.azure.resourcemanager.storage.models.SmbSetting;

/**
 * Samples for FileServices SetServiceProperties.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/FileServicesPut_EnableSecureSmbFeatures.json
     */
    /**
     * Sample code: PutFileServices_EnableSecureSmbFeatures.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        putFileServicesEnableSecureSmbFeatures(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileServices().setServicePropertiesWithResponse("res4410", "sto8607",
            new FileServicePropertiesInner().withProtocolSettings(
                new ProtocolSettings().withSmb(new SmbSetting().withVersions("SMB2.1;SMB3.0;SMB3.1.1")
                    .withAuthenticationMethods("NTLMv2;Kerberos").withKerberosTicketEncryption("RC4-HMAC;AES-256")
                    .withChannelEncryption("AES-128-CCM;AES-128-GCM;AES-256-GCM"))),
            com.azure.core.util.Context.NONE);
    }
}
