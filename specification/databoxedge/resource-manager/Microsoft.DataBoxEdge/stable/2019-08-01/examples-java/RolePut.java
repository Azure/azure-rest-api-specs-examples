
import com.azure.resourcemanager.databoxedge.models.AsymmetricEncryptedSecret;
import com.azure.resourcemanager.databoxedge.models.Authentication;
import com.azure.resourcemanager.databoxedge.models.EncryptionAlgorithm;
import com.azure.resourcemanager.databoxedge.models.IoTDeviceInfo;
import com.azure.resourcemanager.databoxedge.models.IoTRole;
import com.azure.resourcemanager.databoxedge.models.PlatformType;
import com.azure.resourcemanager.databoxedge.models.RoleStatus;
import com.azure.resourcemanager.databoxedge.models.SymmetricKey;
import java.util.Arrays;

/**
 * Samples for Roles CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/RolePut.json
     */
    /**
     * Sample code: RolePut.
     * 
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void rolePut(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager.roles().createOrUpdate("testedgedevice", "IoTRole1", "GroupForEdgeAutomation",
            new IoTRole().withHostPlatform(PlatformType.LINUX).withIoTDeviceDetails(new IoTDeviceInfo()
                .withDeviceId("iotdevice").withIoTHostHub("iothub.azure-devices.net")
                .withAuthentication(new Authentication()
                    .withSymmetricKey(new SymmetricKey().withConnectionString(new AsymmetricEncryptedSecret()
                        .withValue(
                            "Encrypted<<HostName=iothub.azure-devices.net;DeviceId=iotDevice;SharedAccessKey=2C750FscEas3JmQ8Bnui5yQWZPyml0/UiRt1bQwd8=>>")
                        .withEncryptionCertThumbprint("348586569999244")
                        .withEncryptionAlgorithm(EncryptionAlgorithm.AES256)))))
                .withIoTEdgeDeviceDetails(new IoTDeviceInfo().withDeviceId("iotEdge")
                    .withIoTHostHub("iothub.azure-devices.net")
                    .withAuthentication(new Authentication().withSymmetricKey(
                        new SymmetricKey().withConnectionString(new AsymmetricEncryptedSecret().withValue(
                            "Encrypted<<HostName=iothub.azure-devices.net;DeviceId=iotEdge;SharedAccessKey=2C750FscEas3JmQ8Bnui5yQWZPyml0/UiRt1bQwd8=>>")
                            .withEncryptionCertThumbprint("1245475856069999244")
                            .withEncryptionAlgorithm(EncryptionAlgorithm.AES256)))))
                .withShareMappings(Arrays.asList()).withRoleStatus(RoleStatus.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
