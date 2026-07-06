
import com.azure.resourcemanager.storage.fluent.models.FileServicePropertiesInner;
import com.azure.resourcemanager.storage.models.Multichannel;
import com.azure.resourcemanager.storage.models.ProtocolSettings;
import com.azure.resourcemanager.storage.models.SmbSetting;

/**
 * Samples for FileServices SetServiceProperties.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/FileServicesPut_EnableSMBMultichannel.json
     */
    /**
     * Sample code: PutFileServices_EnableSMBMultichannel.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void putFileServicesEnableSMBMultichannel(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getFileServices().setServicePropertiesWithResponse("res4410", "sto8607",
            new FileServicePropertiesInner().withProtocolSettings(new ProtocolSettings()
                .withSmb(new SmbSetting().withMultichannel(new Multichannel().withEnabled(true)))),
            com.azure.core.util.Context.NONE);
    }
}
