import com.azure.resourcemanager.databoxedge.models.AsymmetricEncryptedSecret;
import com.azure.resourcemanager.databoxedge.models.EncryptionAlgorithm;
import com.azure.resourcemanager.databoxedge.models.SecuritySettings;

/** Samples for Devices CreateOrUpdateSecuritySettings. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/SecuritySettingsUpdatePost.json
     */
    /**
     * Sample code: CreateOrUpdateSecuritySettings.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void createOrUpdateSecuritySettings(
        com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .devices()
            .createOrUpdateSecuritySettings(
                "testedgedevice",
                "AzureVM",
                new SecuritySettings()
                    .withDeviceAdminPassword(
                        new AsymmetricEncryptedSecret()
                            .withValue("<value>")
                            .withEncryptionCertThumbprint("7DCBDFC44ED968D232C9A998FC105B5C70E84BE0")
                            .withEncryptionAlgorithm(EncryptionAlgorithm.AES256)),
                com.azure.core.util.Context.NONE);
    }
}
