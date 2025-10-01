
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
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2025-01-01/examples/
     * FileServicesPut_EncryptionInTransitRequired.json
     */
    /**
     * Sample code: PutFileServices_EncryptionInTransitRequired.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        putFileServicesEncryptionInTransitRequired(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getFileServices().setServicePropertiesWithResponse("res4410",
            "sto8607",
            new FileServicePropertiesInner().withProtocolSettings(new ProtocolSettings()
                .withSmb(new SmbSetting().withEncryptionInTransit(new EncryptionInTransit().withRequired(true)))
                .withNfs(new NfsSetting().withEncryptionInTransit(new EncryptionInTransit().withRequired(true)))),
            com.azure.core.util.Context.NONE);
    }
}
