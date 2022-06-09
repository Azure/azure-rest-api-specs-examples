```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.batch.models.Certificate;

/** Samples for Certificate Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/CertificateUpdate.json
     */
    /**
     * Sample code: UpdateCertificate.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void updateCertificate(com.azure.resourcemanager.batch.BatchManager manager) {
        Certificate resource =
            manager
                .certificates()
                .getWithResponse(
                    "default-azurebatch-japaneast",
                    "sampleacct",
                    "sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e",
                    Context.NONE)
                .getValue();
        resource.update().withData("MIIJsgIBAzCCCW4GCSqGSIb3DQE...").withPassword("<ExamplePassword>").apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.
