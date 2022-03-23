Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.batch.models.CertificateFormat;

/** Samples for Certificate Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/CertificateCreate_Full.json
     */
    /**
     * Sample code: CreateCertificate - Full.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void createCertificateFull(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .certificates()
            .define("sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e")
            .withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withData("MIIJsgIBAzCCCW4GCSqGSIb3DQE...")
            .withPassword("<ExamplePassword>")
            .withThumbprintAlgorithm("sha1")
            .withThumbprint("0a0e4f50d51beadeac1d35afc5116098e7902e6e")
            .withFormat(CertificateFormat.PFX)
            .create();
    }
}
```
