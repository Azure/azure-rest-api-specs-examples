```java
/** Samples for Certificate Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/CertificateCreate_Minimal.json
     */
    /**
     * Sample code: CreateCertificate - Minimal Pfx.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void createCertificateMinimalPfx(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .certificates()
            .define("sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e")
            .withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withData("MIIJsgIBAzCCCW4GCSqGSIb3DQE...")
            .withPassword("<ExamplePassword>")
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-batch_1.0.0/sdk/batch/azure-resourcemanager-batch/README.md) on how to add the SDK to your project and authenticate.
