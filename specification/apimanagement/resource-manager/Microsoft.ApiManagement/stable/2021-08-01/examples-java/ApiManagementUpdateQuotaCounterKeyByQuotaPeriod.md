```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.QuotaCounterValueUpdateContract;

/** Samples for QuotaByPeriodKeys Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateQuotaCounterKeyByQuotaPeriod.json
     */
    /**
     * Sample code: ApiManagementUpdateQuotaCounterKeyByQuotaPeriod.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateQuotaCounterKeyByQuotaPeriod(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .quotaByPeriodKeys()
            .updateWithResponse(
                "rg1",
                "apimService1",
                "ba",
                "0_P3Y6M4DT12H30M5S",
                new QuotaCounterValueUpdateContract().withCallsCount(0).withKbTransferred(0.0),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
