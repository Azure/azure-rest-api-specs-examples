```java
import java.util.Arrays;

/** Samples for SingleSignOn CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/SingleSignOn_CreateOrUpdate_MinimumSet_Gen.json
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-dynatrace_1.0.0-beta.1/sdk/dynatrace/azure-resourcemanager-dynatrace/README.md) on how to add the SDK to your project and authenticate.
