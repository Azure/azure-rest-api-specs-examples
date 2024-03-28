
/**
 * Samples for Certificate Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/CertificateCreate_Minimal.json
     */
    /**
     * Sample code: CreateCertificate - Minimal Pfx.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void createCertificateMinimalPfx(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.certificates().define("sha1-0a0e4f50d51beadeac1d35afc5116098e7902e6e")
            .withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct").withPassword("<ExamplePassword>")
            .create();
    }
}
