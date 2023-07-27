import com.azure.resourcemanager.batch.models.CertificateFormat;

/** Samples for Certificate Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/CertificateCreate_Full.json
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
            .withPassword("<ExamplePassword>")
            .withThumbprintAlgorithm("sha1")
            .withThumbprint("0a0e4f50d51beadeac1d35afc5116098e7902e6e")
            .withFormat(CertificateFormat.PFX)
            .create();
    }
}
