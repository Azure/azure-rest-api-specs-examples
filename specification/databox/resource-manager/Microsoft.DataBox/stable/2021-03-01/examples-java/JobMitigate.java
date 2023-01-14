import com.azure.resourcemanager.databox.models.CustomerResolutionCode;
import com.azure.resourcemanager.databox.models.MitigateJobRequest;

/** Samples for ResourceProvider Mitigate. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2021-03-01/examples/JobMitigate.json
     */
    /**
     * Sample code: Mitigate.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void mitigate(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .resourceProviders()
            .mitigateWithResponse(
                "SdkJob8367",
                "SdkRg9836",
                new MitigateJobRequest().withCustomerResolutionCode(CustomerResolutionCode.MOVE_TO_CLEAN_UP_DEVICE),
                com.azure.core.util.Context.NONE);
    }
}
