Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-iothub_1.2.0-beta.1/sdk/iothub/azure-resourcemanager-iothub/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.iothub.models.CertificateVerificationDescription;

/** Samples for Certificates Verify. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_certverify.json
     */
    /**
     * Sample code: Certificates_Verify.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void certificatesVerify(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .certificates()
            .verifyWithResponse(
                "myResourceGroup",
                "myFirstProvisioningService",
                "cert",
                "AAAAAAAADGk=",
                new CertificateVerificationDescription().withCertificate("#####################################"),
                Context.NONE);
    }
}
```
