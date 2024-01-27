
import com.azure.core.util.Context;

/** Samples for MetricAlertsStatus ListByName. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/getMetricAlertStatusByName.
     * json
     */
    /**
     * Sample code: Get an alert rule status.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAnAlertRuleStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getMetricAlertsStatus().listByNameWithResponse("EastUs",
            "custom1",
            "cmVzb3VyY2VJZD0vc3Vic2NyaXB0aW9ucy8xNGRkZjBjNS03N2M1LTRiNTMtODRmNi1lMWZhNDNhZDY4ZjcvcmVzb3VyY2VHcm91cHMvZ2lndGVzdC9wcm92aWRlcnMvTWljcm9zb2Z0LkNvbXB1dGUvdmlydHVhbE1hY2hpbmVzL2dpZ3dhZG1l",
            Context.NONE);
    }
}
