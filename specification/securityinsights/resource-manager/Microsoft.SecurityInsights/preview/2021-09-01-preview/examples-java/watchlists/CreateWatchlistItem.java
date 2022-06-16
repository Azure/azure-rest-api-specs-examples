import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import java.io.IOException;

/** Samples for WatchlistItems CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/watchlists/CreateWatchlistItem.json
     */
    /**
     * Sample code: Creates or updates a watchlist item.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAWatchlistItem(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) throws IOException {
        manager
            .watchlistItems()
            .define("82ba292c-dc97-4dfc-969d-d4dd9e666842")
            .withExistingWatchlist("myRg", "myWorkspace", "highValueAsset")
            .withEtag("0300bf09-0000-0000-0000-5c37296e0000")
            .withItemsKeyValue(
                SerializerFactory
                    .createDefaultManagementSerializerAdapter()
                    .deserialize(
                        "{\"Business tier\":\"10.0.2.0/24\",\"Data tier\":\"10.0.2.0/24\",\"Gateway"
                            + " subnet\":\"10.0.255.224/27\",\"Private DMZ in\":\"10.0.0.0/27\",\"Public DMZ"
                            + " out\":\"10.0.0.96/27\",\"Web Tier\":\"10.0.1.0/24\"}",
                        Object.class,
                        SerializerEncoding.JSON))
            .create();
    }
}
