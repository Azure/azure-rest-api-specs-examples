import com.azure.core.util.Context;

/** Samples for Certificate ListByBatchAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/CertificateListWithFilter.json
     */
    /**
     * Sample code: ListCertificates - Filter and Select.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void listCertificatesFilterAndSelect(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .certificates()
            .listByBatchAccount(
                "default-azurebatch-japaneast",
                "sampleacct",
                null,
                "properties/format,properties/provisioningState",
                "properties/provisioningStateTransitionTime gt '2017-05-01' or properties/provisioningState eq"
                    + " 'Failed'",
                Context.NONE);
    }
}
