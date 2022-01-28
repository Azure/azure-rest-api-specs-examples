Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-deviceprovisioningservices_1.1.0-beta.1/sdk/deviceprovisioningservices/azure-resourcemanager-deviceprovisioningservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for DpsCertificate Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2021-10-15/examples/DPSGetCertificate.json
     */
    /**
     * Sample code: DPSGetCertificate.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSGetCertificate(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .dpsCertificates()
            .getWithResponse("cert", "myResourceGroup", "myFirstProvisioningService", null, Context.NONE);
    }
}
```
