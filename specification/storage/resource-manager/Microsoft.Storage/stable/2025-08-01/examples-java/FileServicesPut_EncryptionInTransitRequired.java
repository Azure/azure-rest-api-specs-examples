
import com.azure.resourcemanager.storage.fluent.models.FileServicePropertiesInner;
import com.azure.resourcemanager.storage.models.EncryptionInTransit;
import com.azure.resourcemanager.storage.models.NfsSetting;
import com.azure.resourcemanager.storage.models.ProtocolSettings;
import com.azure.resourcemanager.storage.models.SmbSetting;

/**
 * Samples for FileServices SetServiceProperties.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/FileServicesPut_EncryptionInTransitRequired.json
     */
    /**
     * Sample code: PutFileServices_EncryptionInTransitRequired.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        putFileServicesEncryptionInTransitRequired(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileServices().setServicePropertiesWithResponse("res4410", "sto8607",
            new FileServicePropertiesInner().withProtocolSettings(new ProtocolSettings()
                .withSmb(new SmbSetting().withEncryptionInTransit(new EncryptionInTransit().withRequired(true)))
                .withNfs(new NfsSetting().withEncryptionInTransit(new EncryptionInTransit().withRequired(true)))),
            com.azure.core.util.Context.NONE);
    }
}
