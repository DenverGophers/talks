# vi: set ft=ruby :

guard 'shell' do

  # anytime a go file changes ...
  watch(%r{.*\.go$}) do

    # announce we're going to build
    system('echo `date "+%H:%M:%S"` - building')

    # build our app; bail if it doesn't go well
    system('go build') || next

    # announce build passed and we're going to test
    system('echo `date "+%H:%M:%S"` - build OK - testing')

    # test our app; bail if it doesn't go well
    system('go test ./...') || next

    # announce test passed and we're going to vet
    system('echo `date "+%H:%M:%S"` - test OK - vetting')

    # vet our app; bail if it doesn't go well
    system('go vet ./...') || next

    # announce vet passed
    system('echo `date "+%H:%M:%S"` - vet OK')

  end

end
