import java.util.UUID;

/** Samples for Recommendations GetGenerateStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/EmptyResponse.json
     */
    /**
     * Sample code: GetGenerateStatus.
     *
     * @param manager Entry point to AdvisorManager.
     */
    public static void getGenerateStatus(com.azure.resourcemanager.advisor.AdvisorManager manager) {
        manager
            .recommendations()
            .getGenerateStatusWithResponse(
                UUID.fromString("00000000-0000-0000-0000-000000000000"), com.azure.core.util.Context.NONE);
    }
}
