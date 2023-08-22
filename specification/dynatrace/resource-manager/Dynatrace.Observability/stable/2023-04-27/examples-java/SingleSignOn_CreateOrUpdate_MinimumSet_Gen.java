import java.util.Arrays;

/** Samples for SingleSignOn CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/SingleSignOn_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: SingleSignOn_CreateOrUpdate_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void singleSignOnCreateOrUpdateMinimumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .singleSignOns()
            .define("default")
            .withExistingMonitor("myResourceGroup", "myMonitor")
            .withSingleSignOnUrl("https://www.dynatrace.io")
            .withAadDomains(Arrays.asList("mpliftrdt20210811outlook.onmicrosoft.com"))
            .create();
    }
}
