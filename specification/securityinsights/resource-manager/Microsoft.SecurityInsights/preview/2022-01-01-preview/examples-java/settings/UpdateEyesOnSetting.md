```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.EyesOn;

/** Samples for ProductSettings Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/settings/UpdateEyesOnSetting.json
     */
    /**
     * Sample code: Update EyesOn settings.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void updateEyesOnSettings(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .productSettings()
            .updateWithResponse(
                "myRg",
                "myWorkspace",
                "EyesOn",
                new EyesOn().withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\""),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
