import com.azure.resourcemanager.dynatrace.models.SingleSignOnStates;
import java.util.Arrays;

/** Samples for SingleSignOn CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/SingleSignOn_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: SingleSignOn_CreateOrUpdate_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void singleSignOnCreateOrUpdateMaximumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .singleSignOns()
            .define("default")
            .withExistingMonitor("myResourceGroup", "myMonitor")
            .withSingleSignOnState(SingleSignOnStates.ENABLE)
            .withEnterpriseAppId("00000000-0000-0000-0000-000000000000")
            .withSingleSignOnUrl("https://www.dynatrace.io")
            .withAadDomains(Arrays.asList("mpliftrdt20210811outlook.onmicrosoft.com"))
            .create();
    }
}
