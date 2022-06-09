```java
import com.azure.core.util.Context;

/** Samples for DpsCertificate Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2021-10-15/examples/DPSDeleteCertificate.json
     */
    /**
     * Sample code: DPSDeleteCertificate.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSDeleteCertificate(
        com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .dpsCertificates()
            .deleteWithResponse(
                "myResourceGroup",
                "AAAAAAAADGk=",
                "myFirstProvisioningService",
                "cert",
                null,
                new byte[0],
                null,
                null,
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-deviceprovisioningservices_1.1.0-beta.1/sdk/deviceprovisioningservices/azure-resourcemanager-deviceprovisioningservices/README.md) on how to add the SDK to your project and authenticate.
