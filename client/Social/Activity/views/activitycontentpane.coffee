class ActivityContentPane extends KDView

  constructor: (options, data) ->
    super options, data

    { lastToFirst, channelId, wrapper, itemClass, typeConstant } = @getOptions()

    @listController = new ActivityListController {
      type          : typeConstant
      viewOptions   :
        itemOptions : { channelId }
      wrapper
      itemClass
      lastToFirst
    }

    @isLoaded = no

  viewAppended: ->
    @addSubView @listController.getView()

  removeItem: (item) ->
    index = @listController.getListView().getItemIndex item
    @listController.removeItem item
    return index

  setContent: (content) ->
    @listController.removeAllItems()
    @listController.instantiateListItems content
    @listController.hideLazyLoader()
    @isLoaded = yes
    @emit 'ContentIsShown'

  appendContent: (content) ->
    @listController.instantiateListItems content
    @listController.hideLazyLoader()

  loadMore: ->
    @emit 'NeedsMoreContent'

  getContentFrom: ->
    @listController.getListView().items.last.getData().createdAt

  getLoadedCount: ->
    @listController.getListView().items.length

  removeMessage: MessagePane::removeMessage

  addItem: (item, index) -> @listController.addItem item, index


